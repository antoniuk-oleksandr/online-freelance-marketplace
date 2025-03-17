<script lang="ts">
  import Icon from '@iconify/svelte'
  import PaginationButton from '../PaginationButton/PaginationButton.svelte'

  type PaginationMoveButtonProps = {
    type: 'prev' | 'next'
    disabled: boolean
    clickAction: () => void
  }

  const { type, ...rest }: PaginationMoveButtonProps = $props()
</script>

<PaginationButton {...rest} currentPage={false}>
  <Icon icon="fe:{type === 'next' ? 'arrow-right' : 'arrow-left'}" />
</PaginationButton>
