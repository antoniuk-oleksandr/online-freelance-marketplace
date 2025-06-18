<script lang="ts" module>
  import OrderByIdReviewTabSuccessModelContentButtons from '../OrderByIdReviewTabSuccessModelContentButtons/OrderByIdReviewTabSuccessModelContentButtons.svelte'
  import OrderByIdReviewTabSuccessModelContentIcon from '../OrderByIdReviewTabSuccessModelContentIcon/OrderByIdReviewTabSuccessModelContentIcon.svelte'
  import OrderByIdReviewTabSuccessModelContentInstructionSection from '../OrderByIdReviewTabSuccessModelContentInstructionSection/OrderByIdReviewTabSuccessModelContentInstructionSection.svelte'
  import OrderByIdReviewTabSuccessModelContentText from '../OrderByIdReviewTabSuccessModelContentText/OrderByIdReviewTabSuccessModelContentText.svelte'
  import OrderByIdReviewTabSuccessModelContentLayout from './OrderByIdReviewTabSuccessModelContentLayout.svelte'

  export { orderByIdReviewTabSuccessModelContent }
</script>

{#snippet orderByIdReviewTabSuccessModelContent()}
  <OrderByIdReviewTabSuccessModelContentLayout>
    <OrderByIdReviewTabSuccessModelContentIcon />
    <OrderByIdReviewTabSuccessModelContentText />
    <OrderByIdReviewTabSuccessModelContentInstructionSection />
    <OrderByIdReviewTabSuccessModelContentButtons />
  </OrderByIdReviewTabSuccessModelContentLayout>
{/snippet}
