export const handleAttachFilesInputChange = (
  e: Event & { currentTarget: EventTarget & HTMLInputElement },
  addFiles: (files: FileList) => void
) => {
  if (!e.target) return;
  const target = e.target as HTMLInputElement;

  if (!target.files) return;
  addFiles(target.files);
}

export const handleAttachFilesButtonClick = (inputRef: HTMLInputElement | undefined) => {
  if (!inputRef) return;
  inputRef.click();
}
