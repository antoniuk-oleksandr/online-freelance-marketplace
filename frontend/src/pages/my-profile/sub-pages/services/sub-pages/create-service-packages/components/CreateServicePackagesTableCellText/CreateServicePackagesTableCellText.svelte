<script lang="ts">
  import type { CreateServicePackagesTableCell } from '@/types/CreateServicePackagesTableCell'

  type CreateServicePackagesTableCellTextProps = {
    cell: CreateServicePackagesTableCell
  }

  const { cell }: CreateServicePackagesTableCellTextProps = $props()
</script>

{#if cell.type === 'text'}
  <p class="p-3 {cell.styles}">{cell.text}</p>
{/if}
