<script lang="ts">
    import CarouselSlideListLayout
        from "@/common-components/ImageCarousel/components/ImageCarouselSlideList/ImageCarouselSlideListLayout.svelte";
    import CarouselSlide from "@/common-components/ImageCarousel/components/ImageCarouselSlide/ImageCarouselSlide.svelte";

    type CarouselSlideListProps = {
        currentIndex: number,
        slides: string[],
    };

    let {slides, currentIndex}: CarouselSlideListProps = $props();
</script>

<CarouselSlideListLayout currentIndex={currentIndex}>
    {#if !slides || slides.length === 0}
        <CarouselSlide slide="/images/no-image.png"/>
    {:else}
        {#each slides as slide}
            <CarouselSlide slide={slide}/>
        {/each}
    {/if}
</CarouselSlideListLayout>
