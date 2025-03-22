<script lang="ts">
  import Avatar from '@/common-components/Avatar/Avatar.svelte'
  import Button from '@/common-components/Button/Button.svelte'
  import StatusElement from '@/common-components/StatusElement/StatusElement.svelte'
  import TableCell from '@/common-components/Table/components/TableCell/TableCell.svelte'
  import TableRow from '@/common-components/Table/components/TableRow/TableRow.svelte'
  import { formatOrderDate } from '@/pages/order-submit-requirements/helpers'
  import type { RequestTableData } from '@/types/RequestTableData'
  import { onMount } from 'svelte'

  type servicesTableBodyProps = {
    requests: RequestTableData[]
  }

  let isLargeScreen = $state<boolean>(false)
  onMount(() => {
    const updateScreenSize = () => (isLargeScreen = window.innerWidth >= 1024)

    updateScreenSize()
    window.addEventListener('resize', updateScreenSize)

    return () => window.removeEventListener('resize', updateScreenSize)
  })

  const { requests }: servicesTableBodyProps = $props()
</script>

{#each requests as request}
  <TableRow
    gridTemplateColumns="grid-cols-1 gap-3 lg:gap-0 lg:grid-cols-requests-table"
    link={isLargeScreen ? `/my-profile/requests/${request.id}` : undefined}
    styles="odd:lg:bg-gray-100 dark:odd:lg:bg-neutral-800 border-light-palette-divider pb-6 mb-6 lg:mb-0 lg:pb-0 dark:border-dark-palette-divider last:pb-0 last:mb-0 last:!border-none border-b"
  >
    <TableCell styles="justify-center lg:justify-start">
      <Avatar
        noSrcImage="/images/no-image.png"
        borderRadius="md"
        src={request.customerAvatar}
        mobileSize="64"
        size="14"
        alt="Customer Avatar"
      />
    </TableCell>

    <TableCell styles="justify-between gap-3" label="Customer">
      <span class="truncate">{request.customerFirstName} {request.customerSurname}</span>
    </TableCell>

    <TableCell styles="justify-between gap-3" label="Service">
      <span class="truncate">{request.title}</span>
    </TableCell>

    <TableCell styles="justify-between" label="Status">
      <StatusElement status={request.status} />
    </TableCell>

    <TableCell styles="justify-between" label="Price">
      <span class="font-semibold">{request.price}$</span>
    </TableCell>

    <TableCell styles="justify-between" label="Date">
      <span>{formatOrderDate(new Date(request.date).toString())}</span>
    </TableCell>

    <TableCell styles="lg:hidden">
      <Button link={`/my-profile/requests/${request.id}`} styles="w-full">View Order</Button>
    </TableCell>
  </TableRow>
{/each}
