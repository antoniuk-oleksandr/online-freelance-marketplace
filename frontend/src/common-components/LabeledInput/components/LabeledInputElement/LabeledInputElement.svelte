<script lang="ts">
  type LabeledInputElementProps = {
    id: string
    placeholder: string
    error?: string | null
    focusAction?: () => void
    blurAction?: () => void
  }

  let props: LabeledInputElementProps = $props()
</script>

<input
  autocomplete="off"
  class="{props.error
    ? '!ring-red-500'
    : ''} placeholder:text-light-palette-text-secondary placeholder:dark:text-dark-palette-text-secondary bg-light-palette-background-block dark:bg-dark-palette-background-block ring-1 ring-light-palette-divider dark:ring-dark-palette-divider text-light-palette-text-primary dark:text-dark-palette-text-primary focus:ring-cyan-500 focus:dark:ring-cyan-500 duration-200 ease-out rounded-md py-2 px-3 outline-none w-full"
  id={props.id}
  name={props.id}
  onblur={props.blurAction}
  onfocus={props.focusAction}
  placeholder={props.placeholder}
  type="text"
/>
