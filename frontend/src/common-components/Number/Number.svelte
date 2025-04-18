<script lang="ts">
  type NumberProps = {
    value: number
  }

  const { value }: NumberProps = $props()
</script>

<div
  class="grid place-items-center font-semibold text-light-palette-background-block dark:text-dark-palette-background-block rounded-full w-6 h-6 bg-light-palette-text-secondary dark:bg-dark-palette-text-secondary"
>
  <span>{value}</span>
</div>
