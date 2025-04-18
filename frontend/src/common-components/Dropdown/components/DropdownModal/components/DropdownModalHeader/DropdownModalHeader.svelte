<script lang="ts">
  import DropdownModalHeaderCloseButton from '../DropdownModalHeaderCloseButton/DropdownModalHeaderCloseButton.svelte'
  import DropdownModalHeaderLayout from './DropdownModalHeaderLayout.svelte'

  type DropdownModalHeaderProps = {
    modalHeaderTitle: string
    setShown?: (value: boolean) => void
  }

  const { modalHeaderTitle, ...rest }: DropdownModalHeaderProps = $props()
</script>

<DropdownModalHeaderLayout>
  <p>{modalHeaderTitle}</p>
  <DropdownModalHeaderCloseButton {...rest} />
</DropdownModalHeaderLayout>
