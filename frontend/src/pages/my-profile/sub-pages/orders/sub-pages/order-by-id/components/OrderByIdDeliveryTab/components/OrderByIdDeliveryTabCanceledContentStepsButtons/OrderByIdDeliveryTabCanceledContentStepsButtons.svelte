<script lang="ts">
  import Button from '@/common-components/Button/Button.svelte'
  import OrderByIdDeliveryTabCanceledContentStepsButtonsLayout from './OrderByIdDeliveryTabCanceledContentStepsButtonsLayout.svelte'
  import Icon from '@iconify/svelte'
</script>

<OrderByIdDeliveryTabCanceledContentStepsButtonsLayout>
  <Button
    linkStyles="!w-full md:!w-fit"
    link="/search"
    color="darkGray"
    styles="!flex gap-2 !w-full md:!w-fit"
  >
    <Icon icon="hugeicons:search-02" height="16" width="16" />
    <span>Browse services</span>
  </Button>
  <Button
    linkStyles="!w-full md:!w-fit"
    styles="!flex gap-2 !w-full md:!w-fit"
    link="/support"
    color="whiteWithGrayOutline"
  >
    <Icon icon="hugeicons:bubble-chat" height="16" width="16" />
    <span>Contact Support</span>
  </Button>
</OrderByIdDeliveryTabCanceledContentStepsButtonsLayout>
