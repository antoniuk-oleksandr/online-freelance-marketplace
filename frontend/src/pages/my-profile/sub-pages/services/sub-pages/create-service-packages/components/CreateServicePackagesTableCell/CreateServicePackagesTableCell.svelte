<script lang="ts">
  import type { CreateServicePackagesTableCell } from '@/types/CreateServicePackagesTableCell'
  import CreateServicePackagesTableCellLayout from './CreateServicePackagesTableCellLayout.svelte'
  import CreateServicePackagesTableCellText from '../CreateServicePackagesTableCellText/CreateServicePackagesTableCellText.svelte'
  import CreateServicePackagesTableCellInput from '../CreateServicePackagesTableCellInput/CreateServicePackagesTableCellInput.svelte'
  import CreateServicePackagesTableCellTextArea from '../CreateServicePackagesTableCellTextArea/CreateServicePackagesTableCellTextArea.svelte'

  type CreateServicePackagesTableCellProps = {
    cell: CreateServicePackagesTableCell
  }

  const { cell }: CreateServicePackagesTableCellProps = $props()
</script>

<CreateServicePackagesTableCellLayout>
  <CreateServicePackagesTableCellText {cell} />
  <CreateServicePackagesTableCellTextArea {cell} />
  <CreateServicePackagesTableCellInput {cell} />
</CreateServicePackagesTableCellLayout>
