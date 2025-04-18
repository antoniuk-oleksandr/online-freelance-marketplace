<script lang="ts">
  import { defineCreateServicePackagesTableCells } from '../../helpers'
  import CreateServicePackagesTableCell from '../CreateServicePackagesTableCell/CreateServicePackagesTableCell.svelte'
  import CreateServicePackagesTableRow from '../CreateServicePackagesTableRow/CreateServicePackagesTableRow.svelte'
  import CreateServicePackagesTableLayout from './CreateServicePackagesTableLayout.svelte'

  const tableCellsData = defineCreateServicePackagesTableCells()
</script>

<CreateServicePackagesTableLayout>
  {#each tableCellsData as row}
    <CreateServicePackagesTableRow>
      {#each row as cell}
        <CreateServicePackagesTableCell {cell} />
      {/each}
    </CreateServicePackagesTableRow>
  {/each}
</CreateServicePackagesTableLayout>
