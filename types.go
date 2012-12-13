package goit

const (
	_GIT_OBJ_ANY       = -2 + iota /**< Object can be any of the following */
	_GIT_OBJ_BAD                   /**< Object is invalid. */
	_GIT_OBJ__EXT1                 /**< Reserved for future use. */
	_GIT_OBJ_COMMIT                /**< A commit object. */
	_GIT_OBJ_TREE                  /**< A tree (directory listing) object. */
	_GIT_OBJ_BLOB                  /**< A file revision object. */
	_GIT_OBJ_TAG                   /**< An annotated tag object. */
	_GIT_OBJ__EXT2                 /**< Reserved for future use. */
	_GIT_OBJ_OFS_DELTA             /**< A delta, base is given by an offset. */
	_GIT_OBJ_REF_DELTA             /**< A delta, base is given by object id. */
)
