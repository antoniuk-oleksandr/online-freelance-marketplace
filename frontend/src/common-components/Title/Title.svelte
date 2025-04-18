<script lang="ts">
  import TitleLayout from './TitleLayout.svelte'

  type TitleProps = {
    text: string
    description?: string
    capitalize?: boolean
  }

  const { text, description, capitalize = true }: TitleProps = $props()
</script>

<TitleLayout>
  <h2
    class="text-light-palette-text-primary dark:text-dark-palette-text-primary text-xl {capitalize ===
    false
      ? ''
      : 'capitalize'} font-semibold"
  >
    {text}
  </h2>
  {#if description}
    <p class="text-light-palette-text-secondary dark:text-dark-palette-text-secondary text-base">
      {description}
    </p>
  {/if}
</TitleLayout>
