<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'
  import DropdownModalHeader from './components/DropdownModalHeader/DropdownModalHeader.svelte'
  import DropdownModalLayout from './DropdownModalLayout.svelte'

  type DropdownModalProps = LayoutProps & {
    modalHeaderTitle: string
  }

  const { children, ...rest }: DropdownModalProps = $props()
</script>

<DropdownModalLayout>
  <DropdownModalHeader {...rest} />
  {@render children()}
</DropdownModalLayout>
