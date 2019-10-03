package dockertest

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"go.uber.org/zap"
)

type Pool struct {
	logger *zap.Logger
	cli    *Client
}

func NewPool(
	logger *zap.Logger,
	cli *Client,
) *Pool {
	logger.Info("creating dockertest pool")
	pool := &Pool{
		logger: logger,
		cli:    cli,
	}
	logger.Info("created dockertest pool")
	return pool
}

func (p *Pool) PullOfficialImage(ctx context.Context, imageName string, tag string) error {
	canonicalName := fmt.Sprintf("docker.io/library/%s:%s", imageName, tag)
	logger := p.logger.With(zap.String("canonicalName", canonicalName))
	logger.Info("pulling image")
	reader, err := p.cli.ImagePull(ctx, canonicalName, types.ImagePullOptions{})
	if err != nil {
		return fmt.Errorf("pull official image: %w", err)
	}
	_, err = io.Copy(ioutil.Discard, reader)
	if err != nil {
		return fmt.Errorf("pull official image: %w", err)
	}
	logger.Info("pulled image")
	return nil
}

func (p *Pool) StartContainer(ctx context.Context, imageName string, tag string) (string, error) {
	imageLogger := p.logger.With(
		zap.String("imageName", imageName),
		zap.String("tag", tag),
	)
	imageLogger.Info(
		"pulling image of container to run",
		zap.String("imageName", imageName),
		zap.String("tag", tag),
	)
	if err := p.PullOfficialImage(ctx, imageName, tag); err != nil {
		return "", fmt.Errorf("start container: %w", err)
	}
	imageLogger.Info("image pulled, creating container")
	containerCfg := &container.Config{
		Image: fmt.Sprintf("%v:%v", imageName, tag),
	}
	createResponse, err := p.cli.ContainerCreate(ctx, containerCfg, nil, nil, "")
	if err != nil {
		return "", fmt.Errorf("start container: %w", err)
	}
	id := createResponse.ID
	idLogger := p.logger.With(zap.String("id", id))
	idLogger.Info(
		"created container, now starting it",
	)
	if err := p.cli.ContainerStart(ctx, id, types.ContainerStartOptions{}); err != nil {
		return "", fmt.Errorf("start container: %w", err)
	}
	idLogger.Info("container started")
	return id, nil
}

func (p *Pool) StopContainer(ctx context.Context, containerID string, timeout time.Duration) error {
	idLogger := p.logger.With(zap.String("id", containerID))
	idLogger.Info("stopping container")
	if err := p.cli.ContainerStop(ctx, containerID, &timeout); err != nil {
		return fmt.Errorf("stop container: %w", err)
	}
	idLogger.Info("container stopped, now removing it")
	if err := p.cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{}); err != nil {
		return fmt.Errorf("stop container: %w", err)
	}
	idLogger.Info("container removed")
	return nil
}