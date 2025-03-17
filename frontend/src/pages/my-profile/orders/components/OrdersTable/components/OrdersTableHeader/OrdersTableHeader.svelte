<script lang="ts">
  import TableCell from '@/common-components/Table/components/TableCell/TableCell.svelte'
  import TableRow from '@/common-components/Table/components/TableRow/TableRow.svelte'
</script>

<TableRow
  hiddenOnSmall={true}
  gridTemplateColumns="grid-cols-orders-table"
  styles="font-semibold bg-gray-100 dark:bg-neutral-800 bg-p"
>
  <TableCell>Image</TableCell>
  <TableCell>Order ID</TableCell>
  <TableCell>Title</TableCell>
  <TableCell>Status</TableCell>
  <TableCell>Price</TableCell>
  <TableCell>Date</TableCell>
</TableRow>
