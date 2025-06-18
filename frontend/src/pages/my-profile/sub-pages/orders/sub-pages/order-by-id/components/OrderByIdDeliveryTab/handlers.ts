import { downloadFile } from "./helpers"

export const handleDownloadButtonClick = async (file: string | string[]) => {
  if (Array.isArray(file)) {
    await Promise.all(
      file.map(async (item) => await downloadFile(item))
    )
  } else await downloadFile(file);
}
