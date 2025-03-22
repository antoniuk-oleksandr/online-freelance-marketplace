<script lang="ts">
  import TableCell from '@/common-components/Table/components/TableCell/TableCell.svelte'
  import TableRow from '@/common-components/Table/components/TableRow/TableRow.svelte'
</script>

<TableRow
  hiddenOnSmall={true}
  gridTemplateColumns="grid-cols-requests-table"
  styles="font-semibold bg-gray-100 dark:bg-neutral-800"
>
  <TableCell>Customer Avatar</TableCell>
  <TableCell>Customer Name</TableCell>
  <TableCell>Service Title</TableCell>
  <TableCell>Status</TableCell>
  <TableCell>Price</TableCell>
  <TableCell>Date</TableCell>
</TableRow>
