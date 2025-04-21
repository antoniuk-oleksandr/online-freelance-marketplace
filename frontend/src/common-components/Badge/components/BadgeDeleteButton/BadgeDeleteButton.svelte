<script lang="ts">
  import Icon from '@iconify/svelte'
  import BadgeDeleteButtonLayout from './BadgeDeleteButtonLayout.svelte'

  type BadgeDeleteButtonProps = {
    deleteFunc: (index: number) => void
    index: number
    isMobile: boolean
  }

  const props: BadgeDeleteButtonProps = $props()
</script>

<BadgeDeleteButtonLayout {...props}>
  <Icon width={props.isMobile ? 16 : 14} height={props.isMobile ? 16 : 14} icon="ci:close-lg" />
</BadgeDeleteButtonLayout>
