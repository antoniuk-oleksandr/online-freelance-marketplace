<script lang="ts">
    import CarouselControlLayout
        from "@/common-components/ImageCarousel/components/ImageCarouselControl/ImageCarouselControlLayout.svelte";
    import CarouselControlIcon
        from "@/common-components/ImageCarousel/components/ImageCarouselControlIcon/ImageCarouselControlIcon.svelte";

    type CarouselControlProps = {
        nextSlide: () => void;
        prevSlide: () => void;
    };

    const {nextSlide, prevSlide}: CarouselControlProps = $props();
</script>

<CarouselControlLayout>
    <CarouselControlIcon clickAction={prevSlide} direction="left"/>
    <CarouselControlIcon clickAction={nextSlide} direction="right"/>
</CarouselControlLayout>