<script lang="ts">
    import HomeSectionLayout from "@/pages/home/common-components/HomeSectionLayout/HomeSectionLayout.svelte";
    import HomeSectionHeader from "@/pages/home/common-components/HomeSectionHeader/HomeSectionHeader.svelte";
    import HomeReviewsSectionItemList
        from "@/pages/home/components/HomeReviewsSection/components/HomeReviewsSectionItemList/HomeReviewsSectionItemList.svelte";
    import type {BestReview} from "@/types/BestReview.ts";

    type HomeReviewsSectionProps = {
        bestReviews: BestReview[];
    };

    const props: HomeReviewsSectionProps = $props();
</script>

<HomeSectionLayout>
    <HomeSectionHeader
            title="What Our Clients Say"
            description="Read reviews from satisfied customers"
    />
    <HomeReviewsSectionItemList {...props}/>
</HomeSectionLayout>
