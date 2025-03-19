<script lang="ts">
  import TableCell from '@/common-components/Table/components/TableCell/TableCell.svelte'
  import TableRow from '@/common-components/Table/components/TableRow/TableRow.svelte'
</script>

<TableRow
  hiddenOnSmall={true}
  gridTemplateColumns="grid-cols-services-table"
  styles="font-semibold bg-gray-100 dark:bg-neutral-800 bg-p"
>
  <TableCell>Image</TableCell>
  <TableCell>Service ID</TableCell>
  <TableCell>Title</TableCell>
  <TableCell>Category</TableCell>
  <TableCell>Rating</TableCell>
  <TableCell>Orders Count</TableCell>
  <TableCell>Price</TableCell>
  <TableCell>Date</TableCell>
</TableRow>
