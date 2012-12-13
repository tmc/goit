package goit

const (
	GIT_OBJ_ANY       = -2 + iota /**< Object can be any of the following */
	GIT_OBJ_BAD                   /**< Object is invalid. */
	_GIT_OBJ__EXT1                /**< Reserved for future use. */
	GIT_OBJ_COMMIT                /**< A commit object. */
	GIT_OBJ_TREE                  /**< A tree (directory listing) object. */
	GIT_OBJ_BLOB                  /**< A file revision object. */
	GIT_OBJ_TAG                   /**< An annotated tag object. */
	_GIT_OBJ__EXT2                /**< Reserved for future use. */
	GIT_OBJ_OFS_DELTA             /**< A delta, base is given by an offset. */
	GIT_OBJ_REF_DELTA             /**< A delta, base is given by object id. */
)
