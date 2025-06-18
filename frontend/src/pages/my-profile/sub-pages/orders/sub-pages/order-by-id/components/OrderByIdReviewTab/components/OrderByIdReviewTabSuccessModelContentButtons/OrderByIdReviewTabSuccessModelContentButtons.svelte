<script lang="ts">
  import Button from '@/common-components/Button/Button.svelte'
  import { modalStore } from '@/common-stores/modal-store'
  import { navigate } from 'svelte-routing'
  import OrderByIdReviewTabSuccessModelContentButtonsLayout from './OrderByIdReviewTabSuccessModelContentButtonsLayout.svelte'

  const closeModal = () => {
    modalStore.set({ isOpened: false, renderContent: undefined, title: undefined })
  }
</script>

<OrderByIdReviewTabSuccessModelContentButtonsLayout>
  <Button
    clickAction={closeModal}
    styles="!h-12 !basis-full md:!basis-1/2"
    color="whiteWithGrayOutline">View Your Review</Button
  >
  <Button
    clickAction={() => {
      closeModal()
      navigate('/search')
    }}
    styles="!h-12 !basis-full md:!basis-1/2"
    color="darkGray">Continue Browsing</Button
  >
</OrderByIdReviewTabSuccessModelContentButtonsLayout>
