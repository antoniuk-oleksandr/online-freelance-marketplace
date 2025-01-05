<script lang="ts">
    import CarouselSlideListLayout
        from "@/common-components/ImageCarousel/components/ImageCarouselSlideList/ImageCarouselSlideListLayout.svelte";
    import CarouselSlide
        from "@/common-components/ImageCarousel/components/ImageCarouselSlide/ImageCarouselSlide.svelte";
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    type CarouselSlideListProps = LayoutProps & {
        currentIndex: number,
    };

    let {children, currentIndex}: CarouselSlideListProps = $props();

</script>

<CarouselSlideListLayout currentIndex={currentIndex}>
    {@render children()}
</CarouselSlideListLayout>