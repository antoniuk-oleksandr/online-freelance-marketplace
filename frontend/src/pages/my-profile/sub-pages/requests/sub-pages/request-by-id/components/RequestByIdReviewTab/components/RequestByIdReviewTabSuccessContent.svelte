<script lang="ts">
  import Avatar from '@/common-components/Avatar/Avatar.svelte'
  import Label from '@/common-components/Label/Label.svelte'
  import Rating from '@/common-components/Rating/Rating.svelte'
  import Title from '@/common-components/Title/Title.svelte'
  import OrderByIdDeliveryTabContentMessage from '@/pages/my-profile/sub-pages/orders/sub-pages/order-by-id/components/OrderByIdDeliveryTab/components/OrderByIdDeliveryTabContentMessage/OrderByIdDeliveryTabContentMessage.svelte'
  import type { OrderReview } from '@/types/OrderReview'
  import { Link } from 'svelte-routing'

  type Props = {
    data: OrderReview
  }

  const { data }: Props = $props()
</script>

<div class="flex flex-col w-full h-full gap-6">
  <Title description="Feedback from the customer about the order" text="Customer Review" />
  <div class="flex gap-3 items-center">
    <Avatar size="24" alt="avatar" src={data.customer.avatar} />
    <div>
      <p
        class="text-xl font-semibold text-light-palette-text-primary dark:text-dark-palette-text-primary"
      >
        {data.customer.firstName}
        {data.customer.surname}
      </p>
      <Link to="/users/{data.customer.id}">
        <span class="text-lg font-medium hover:underline">@{data.customer.username}</span>
      </Link>
    </div>
  </div>
  <OrderByIdDeliveryTabContentMessage title="Review message" message={data.content} />
  <div class="flex flex-col gap-1">
    <Label text="Rating" />
    <div class="flex  gap-3 items-center justify-center">
      <Rating iconStyles="!text-[32px]" value={data.rating} />
      <p
        class="w-16 text-lg font-semibold text-light-palette-text-primary dark:text-dark-palette-text-primary"
      >
        {data.rating}.0/5.0
      </p>
    </div>
  </div>
</div>
