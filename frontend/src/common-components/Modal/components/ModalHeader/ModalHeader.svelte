<script lang="ts">
  import ModalHeaderLayout from '@/common-components/Modal/components/ModalHeader/ModalHeaderLayout.svelte'
  import ModalHeaderCloseButton from '@/common-components/Modal/components/ModalHeaderCloseButton/ModalHeaderCloseButton.svelte'

  type ModalHeaderProps = {
    title?: string
    headerStyles?: string
  }

  const { title, headerStyles }: ModalHeaderProps = $props()
</script>

{#if title}
  <ModalHeaderLayout {headerStyles}>
    <span>{title}</span>
    <ModalHeaderCloseButton />
  </ModalHeaderLayout>
{/if}
