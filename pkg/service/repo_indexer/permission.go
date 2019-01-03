// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// Auto generated by 'go run gen_helper.go', DO NOT EDIT.

package repo_indexer

import (
	"context"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/db"
	"openpitrix.io/openpitrix/pkg/gerr"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pi"
	"openpitrix.io/openpitrix/pkg/util/ctxutil"
)

func CheckRepoEventsPermission(ctx context.Context, resourceIds []string) ([]*models.RepoEvent, error) {
	if len(resourceIds) == 0 {
		return nil, nil
	}
	var sender = ctxutil.GetSender(ctx)
	var repoevents []*models.RepoEvent
	_, err := pi.Global().DB(ctx).
		Select(models.RepoEventColumns...).
		From(constants.TableRepoEvent).
		Where(db.Eq(constants.ColumnRepoEventId, resourceIds)).Load(&repoevents)
	if err != nil {
		return nil, gerr.NewWithDetail(ctx, gerr.Internal, err, gerr.ErrorInternalError)
	}
	if sender != nil {
		for _, repoevent := range repoevents {
			if !repoevent.OwnerPath.CheckPermission(sender) {
				return nil, gerr.New(ctx, gerr.PermissionDenied, gerr.ErrorResourceAccessDenied, repoevent.RepoEventId)
			}
		}
	}
	if len(repoevents) == 0 {
		return nil, gerr.New(ctx, gerr.NotFound, gerr.ErrorResourceNotFound, resourceIds)
	}
	return repoevents, nil
}

func CheckRepoEventPermission(ctx context.Context, resourceId string) (*models.RepoEvent, error) {
	if len(resourceId) == 0 {
		return nil, nil
	}
	var sender = ctxutil.GetSender(ctx)
	var repoevents []*models.RepoEvent
	_, err := pi.Global().DB(ctx).
		Select(models.RepoEventColumns...).
		From(constants.TableRepoEvent).
		Where(db.Eq(constants.ColumnRepoEventId, resourceId)).Load(&repoevents)
	if err != nil {
		return nil, gerr.NewWithDetail(ctx, gerr.Internal, err, gerr.ErrorInternalError)
	}
	if sender != nil {
		for _, repoevent := range repoevents {
			if !repoevent.OwnerPath.CheckPermission(sender) {
				return nil, gerr.New(ctx, gerr.PermissionDenied, gerr.ErrorResourceAccessDenied, repoevent.RepoEventId)
			}
		}
	}
	if len(repoevents) == 0 {
		return nil, gerr.New(ctx, gerr.NotFound, gerr.ErrorResourceNotFound, resourceId)
	}
	return repoevents[0], nil
}
