<script lang="ts">
  import type { OrderTableData } from '@/types/OrderTableData'
  import Avatar from '@/common-components/Avatar/Avatar.svelte'
  import StatusElement from '@/common-components/StatusElement/StatusElement.svelte'
  import TableCell from '@/common-components/Table/components/TableCell/TableCell.svelte'
  import TableRow from '@/common-components/Table/components/TableRow/TableRow.svelte'
  import { formatOrderDate } from '@/pages/order-submit-requirements/helpers'
  import Button from '@/common-components/Button/Button.svelte'
  import { onMount } from 'svelte'
  import { defineOrderTableLink, defineOrderTableButtonText } from '../../../../helpers'

  type OrdersTableBodyProps = {
    orders: OrderTableData[]
  }

  let isLargeScreen = $state<boolean>(false)
  onMount(() => {
    const updateScreenSize = () => (isLargeScreen = window.innerWidth >= 1024)

    updateScreenSize()
    window.addEventListener('resize', updateScreenSize)

    return () => window.removeEventListener('resize', updateScreenSize)
  })

  const { orders }: OrdersTableBodyProps = $props()
</script>

{#each orders as order}
  <TableRow
    gridTemplateColumns="grid-cols-1 gap-3 lg:gap-0 lg:grid-cols-orders-table"
    link={defineOrderTableLink(order, isLargeScreen)}
    styles="odd:lg:bg-gray-100 dark:odd:lg:bg-neutral-800 border-light-palette-divider pb-6 mb-6 lg:mb-0 lg:pb-0 dark:border-dark-palette-divider last:pb-0 last:mb-0 last:!border-none border-b"
  >
    <TableCell styles="justify-center">
      <Avatar
        noSrcImage="/images/no-image.png"
        borderRadius="md"
        src={order.image}
        mobileSize="64"
        size="14"
        alt="Order Image"
      />
    </TableCell>
    <TableCell styles="justify-between gap-3" label="Order ID">
      <span>#{order.id}</span>
    </TableCell>
    <TableCell styles="justify-between gap-3 items-center " label="Title">
      <span class="truncate">{order.title}</span>
    </TableCell>
    <TableCell styles="justify-between" label="Status">
      <StatusElement status={order.status} />
    </TableCell>
    <TableCell styles="justify-between" label="Price">
      <span class="font-semibold">{order.price}$</span>
    </TableCell>
    <TableCell styles="justify-between" label="Date">
      <span>{formatOrderDate(new Date(order.date).toString())}</span>
    </TableCell>
    <TableCell styles="lg:hidden">
      <Button link={defineOrderTableLink(order)} styles="w-full"
        >{defineOrderTableButtonText(order)}
      </Button>
    </TableCell>
  </TableRow>
{/each}
