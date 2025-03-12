<script lang="ts">
  type InputLabelProps = {
    id: string
    value: string
    label: string
    error: string | null
    wasSubmitted: boolean
    isFocused: boolean
  }

  const { id, value, label, wasSubmitted, isFocused, error }: InputLabelProps = $props()
</script>

<label
  for={id}
  class="{isFocused || value.length > 0
    ? 'text-cyan-500 -translate-y-[32.5px] text-sm top-[23px]'
    : 'text-base top-[12px] text-light-palette-divider dark:text-dark-palette-divider'}
                {wasSubmitted && error ? 'text-red-500' : ''}
                duration-translate cursor-text bg-light-palette-background-default dark:bg-dark-palette-background-default lg:bg-light-palette-background-block lg:dark:bg-dark-palette-background-block px-1  absolute left-3"
  >{label}</label
>
