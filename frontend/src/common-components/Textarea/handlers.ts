export const handleTextareaInput = (
  e: Event, setValue?: (value: string) => void

) => {
  if (!setValue) return;
  const target = e.target as HTMLTextAreaElement;
  const value = target.value;

  setValue(value);
}
