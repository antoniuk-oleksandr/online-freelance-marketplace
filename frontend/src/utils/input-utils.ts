import type { EventHandler } from "svelte/elements";

export const allowOnlyFloatInput: EventHandler<InputEvent, HTMLInputElement> = (event) => {
  const input = event.target as HTMLInputElement;

  if (!input) return;

  const { selectionStart, selectionEnd, value } = input;
  const data = event.data ?? '';

  const nextValue =
    value.substring(0, selectionStart || 0) +
    data +
    value.substring(selectionEnd || 0);

  const valid = nextValue === '' || /^(\d+)?(\.\d*)?$/.test(nextValue);

  if (!valid) {
    event.preventDefault();
  }
};
