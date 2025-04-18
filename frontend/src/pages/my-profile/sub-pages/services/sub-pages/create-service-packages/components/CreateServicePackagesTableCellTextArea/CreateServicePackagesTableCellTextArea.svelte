<script lang="ts">
  import type { CreateServicePackagesTableCell } from '@/types/CreateServicePackagesTableCell'
  import CreateServicePackagesTableCellTextAreaLayout from './CreateServicePackagesTableCellTextAreaLayout.svelte'
  import Icon from '@iconify/svelte'

  type CreateServicePackagesTableTextAreaProps = {
    cell: CreateServicePackagesTableCell
  }

  const { cell }: CreateServicePackagesTableTextAreaProps = $props()
</script>

<CreateServicePackagesTableCellTextAreaLayout>
  {#if cell.type === 'textarea'}
    <!-- svelte-ignore element_invalid_self_closing_tag -->
    <textarea
      rows="5"
      name={cell.name}
      class="peer w-full p-3 bg-transparent resize-none outline-none placeholder:text-light-palette-text-secondary placeholder:dark:text-dark-palette-text-secondary"
      placeholder={cell.placeholder}
    />

    {#key cell.rightIcon}
      <Icon
        class="transition-opacity ease-out duration-200 absolute right-3 top-3 opacity-100 peer-focus:opacity-0 pointer-events-none"
        width={24}
        height={24}
        icon="hugeicons:pencil-edit-01"
      />
    {/key}
  {/if}
</CreateServicePackagesTableCellTextAreaLayout>
