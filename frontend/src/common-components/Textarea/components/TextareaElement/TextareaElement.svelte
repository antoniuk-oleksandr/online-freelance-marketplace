<script lang="ts">
  type TextareaElementProps = {
    placeholder: string
    id: string
    name: string
    error: string | undefined
    setFocus: (value: boolean) => void
    styles?: string
    useCounter: boolean
    wasSubmitted: boolean
  }

  const { styles, wasSubmitted, error, setFocus, useCounter, ...rest }: TextareaElementProps =
    $props()
</script>

<!-- svelte-ignore element_invalid_self_closing_tag -->
<textarea
  onfocus={() => setFocus(true)}
  onblur={() => setFocus(false)}
  rows="4"
  {...rest}
  class="{styles} {error && wasSubmitted ? 'ring-red-500 placeholder:text-red-500' : ''} {useCounter
    ? 'rounded-t-xl'
    : 'rounded-xl'} placeholder:text-light-palette-text-secondary placeholder:dark:text-dark-palette-text-secondary w-full p-6 bg-light-palette-background-block dark:bg-dark-palette-background-block text-base rounded-md outline-none resize-none"
/>
