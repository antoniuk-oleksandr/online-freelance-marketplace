<script lang="ts">
    import {handleMoreReviewsButtonClick} from "@/pages/users/handlers.ts";
    import Button from "@/common-components/Button/Button.svelte";

    type ShowMoreReviewsButtonProps = {
        showMoreReviewsButtonAction: () => Promise<void>,
        hasMore: boolean,
        buttonSuffix: string,
    }

    const {showMoreReviewsButtonAction, hasMore, buttonSuffix}: ShowMoreReviewsButtonProps = $props();

    let loading = $state<boolean>(false);
    const setLoading = (value: boolean) => loading = value;
</script>

{#if hasMore}
    <Button
            loading={loading}
            clickAction={() => handleMoreReviewsButtonClick(setLoading, showMoreReviewsButtonAction)}
            styles="text-base w-48 capitalize"
            color="outline"
    >show more {buttonSuffix}
    </Button>
{/if}