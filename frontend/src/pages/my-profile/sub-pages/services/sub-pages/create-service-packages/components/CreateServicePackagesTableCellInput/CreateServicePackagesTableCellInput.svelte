<script lang="ts">
  import type { CreateServicePackagesTableCell } from '@/types/CreateServicePackagesTableCell'
  import CreateServicePackagesTableCellInputLayout from './CreateServicePackagesTableCellInputLayout.svelte'
  import { allowOnlyFloatInput } from '@/utils/input-utils'

  type CreateServicePackagesTableCellInputProps = {
    cell: CreateServicePackagesTableCell
  }

  const { cell }: CreateServicePackagesTableCellInputProps = $props()
</script>

{#if cell.type === 'input'}
  <CreateServicePackagesTableCellInputLayout>
    <input
      name={cell.name}
      type={cell.inputType === 'number' ? 'number' : 'text'}
      class="{cell.styles} w-full p-3 bg-transparent outline-none placeholder:text-light-palette-text-secondary placeholder:dark:text-dark-palette-text-secondary"
      placeholder={cell.placeholder}
      onbeforeinput={(e) => cell.inputType === 'number' && allowOnlyFloatInput(e)}
    />
    {#key cell.rightBadge}
      <span>{cell.rightBadge}</span>
    {/key}
  </CreateServicePackagesTableCellInputLayout>
{/if}
