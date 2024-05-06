package git

import (
	"context"
	"fmt"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/pkg/errors"
)

// FetchSource fetches source from git repository and checkouts to given hash.
// You can access to source using fs.
func FetchSource(ctx context.Context, fs billy.Filesystem, repo, hash string) error {
	dotgit, err := fs.Chroot(git.GitDirName)
	if err != nil {
		return errors.Wrap(err, "changing root dir to .git")
	}

	s := filesystem.NewStorage(dotgit, cache.NewObjectLRUDefault())

	r, err := clone(ctx, fs, s, repo)
	if err != nil {
		return errors.Wrap(err, "cloning repo")
	}

	wt, err := r.Worktree()
	if err != nil {
		return errors.Wrap(err, "getting worktree")
	}

	return wt.Checkout(&git.CheckoutOptions{Hash: plumbing.NewHash(hash)})
}

// clone performs platform-dependent cloning via GitHub.
func clone(ctx context.Context, fs billy.Filesystem, s storage.Storer, path string) (*git.Repository, error) {
	return git.CloneContext(ctx, s, fs, &git.CloneOptions{
		URL: fmt.Sprintf("https://github.com/%s.git", path),
	})
}
