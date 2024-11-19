<script lang="ts">
    import CarouselControlLayout
        from "@/common-components/Carousel/components/CarouselControl/CarouselControlLayout.svelte";
    import CarouselControlIcon
        from "@/common-components/Carousel/components/CarouselControlIcon/CarouselControlIcon.svelte";

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