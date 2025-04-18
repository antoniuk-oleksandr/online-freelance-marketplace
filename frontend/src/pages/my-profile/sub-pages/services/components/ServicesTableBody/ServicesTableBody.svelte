<script lang="ts">
  import Avatar from '@/common-components/Avatar/Avatar.svelte'
  import Button from '@/common-components/Button/Button.svelte'
  import TableCell from '@/common-components/Table/components/TableCell/TableCell.svelte'
  import TableRow from '@/common-components/Table/components/TableRow/TableRow.svelte'
  import { formatOrderDate } from '@/pages/order-submit-requirements/helpers'
  import type { ServiceTableData } from '@/types/ServiceTableData'
  import { onMount } from 'svelte'

  type servicesTableBodyProps = {
    services: ServiceTableData[]
  }

  let isLargeScreen = $state<boolean>(false)
  onMount(() => {
    const updateScreenSize = () => (isLargeScreen = window.innerWidth >= 1024)

    updateScreenSize()
    window.addEventListener('resize', updateScreenSize)

    return () => window.removeEventListener('resize', updateScreenSize)
  })

  const { services }: servicesTableBodyProps = $props()
</script>

{#each services as service}
  <TableRow
    gridTemplateColumns="grid-cols-1 gap-3 lg:gap-0 lg:grid-cols-services-table"
    link={isLargeScreen ? `/my-profile/services/${service.id}` : undefined}
    styles="odd:lg:bg-gray-100 dark:odd:lg:bg-neutral-800 border-light-palette-divider pb-6 mb-6 lg:mb-0 lg:pb-0 dark:border-dark-palette-divider last:pb-0 last:mb-0 last:!border-none border-b"
  >
    <TableCell styles="justify-center">
      <Avatar
        noSrcImage="/images/no-image.png"
        borderRadius="md"
        src={service.image}
        mobileSize="64"
        size="14"
        alt="service Image"
      />
    </TableCell>
    <TableCell styles="justify-between gap-3" label="Service ID">
      <span>#{service.id}</span>
    </TableCell>
    <TableCell styles="justify-between gap-3 items-center " label="Title">
      <span class="truncate">{service.title}</span>
    </TableCell>
    <TableCell styles="justify-between gap-3 items-center " label="Category">
      <span class="truncate">{service.category}</span>
    </TableCell>
    <TableCell styles="justify-between" label="Rating">
      <span>{service.rating === 0 ? '-' : service.rating}</span>
    </TableCell>
    <TableCell styles="justify-between" label="Orders Count">
      <span>{service.ordersCount === 0 ? '-' : service.ordersCount}</span>
    </TableCell>
    <TableCell styles="justify-between" label="Price">
      <span class="font-semibold">{service.price}$</span>
    </TableCell>
    <TableCell styles="justify-between" label="Date">
      <span>{formatOrderDate(new Date(service.date).toString())}</span>
    </TableCell>
    <TableCell styles="lg:hidden">
      <Button link="/my-profile/services/{service.id}" styles="w-full">View service</Button>
    </TableCell>
  </TableRow>
{/each}
