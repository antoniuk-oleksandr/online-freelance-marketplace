<script lang="ts">
    import CarouselSlideListLayout
        from "@/common-components/Carousel/components/CarouselSlideList/CarouselSlideListLayout.svelte";
    import CarouselSlide from "@/common-components/Carousel/components/CarouselSlide/CarouselSlide.svelte";

    type CarouselSlideListProps = {
        currentIndex: number,
        slides: string[],
    };

    let {slides, currentIndex}: CarouselSlideListProps = $props();
</script>

<CarouselSlideListLayout currentIndex={currentIndex}>
    {#if !slides || slides.length === 0}
        <CarouselSlide slide="/images/no-image.jpg"/>
    {:else}
        {#each slides as slide}
            <CarouselSlide slide={slide}/>
        {/each}
    {/if}
</CarouselSlideListLayout>