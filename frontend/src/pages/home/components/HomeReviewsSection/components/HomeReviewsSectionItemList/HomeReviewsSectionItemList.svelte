<script lang="ts">
    import HomeReviewsSectionItemListLayout
        from "@/pages/home/components/HomeReviewsSection/components/HomeReviewsSectionItemList/HomeReviewsSectionItemListLayout.svelte";
    import HomeReviewsSectionItem
        from "@/pages/home/components/HomeReviewsSection/components/HomeReviewsSectionItem/HomeReviewsSectionItem.svelte";
    import type {BestReview} from "@/types/BestReview.js";

    type HomeReviewsSectionItemListProps = {
        bestReviews: BestReview[];
    };

    const {bestReviews}: HomeReviewsSectionItemListProps = $props();
</script>

<HomeReviewsSectionItemListLayout>
    {#each bestReviews as review}
        <HomeReviewsSectionItem {...review}/>
    {/each}
</HomeReviewsSectionItemListLayout>