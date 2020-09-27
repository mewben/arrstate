export const FILE_THUMBNAIL_SIZE = {
  SMALL: "100x100",
  MEDIUM: "400x400",
}

export const FILE_FORMAT_TYPE = {
  IMAGE: "image",
  VIDEO: "video",
  DOCUMENT: "document",
  APPLICATION: "application",
  LINK: "link",
  AUDIO: "audio",
  TEXT: "text",
}

export const FILE_ICONS = {
  FILE: {
    GENERIC: "file",
    IMAGE: {},
    VIDEO: {},
    AUDIO: {},
    TEXT: {},
    APPLICATION: {},
  },
  DOCUMENT: {
    GENERIC: "document",
  },
  LINK: {
    GENERIC: "link",
  },
}

export const FILE_UPLOAD_S3_FOLDERS = {
  PENDING: "pending",
  PROCESSED: "processed/s",
  DELETED: "deleted",
}

export const FILE_UPLOAD_SETTINGS = {
  MAX_SIZE: 20 * 1024 * 1024, // 20MB
  MAX_FILES: 20,
}

export const FILE_TYPE = {
  FILE: "file",
  DOCUMENT: "document",
  LINK: "link",
}

export const FILE_EXTENSION_TYPES = {
  IMAGE: ["jpg", "png", "gif"],
  VIDEO: [],
}

export const FILE_PREVIEW_EXTENSIONS = [
  "pdf",
  "ppt",
  "pptx",
  "xls",
  "xlsx",
  "doc",
  "docx",
  "odt",
  "odp",
  "ai",
  "psd",
]

export const FILE_ACCEPT = {
  ANY_IMAGE: ["image/*"],
  REAL_IMAGE: ["image/jpeg", "image/gif", "image/png"],
}

export const FILE_DIMENSIONS = {
  USER_AVATAR: [200, 200],
  BUSINESS_LOGO: [400, 400],
  FAVICON: [128, 128],
  PROFILE_COVER: [400, 400],
}
