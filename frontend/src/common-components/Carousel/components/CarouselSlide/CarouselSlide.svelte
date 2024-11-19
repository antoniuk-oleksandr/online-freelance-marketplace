<script lang="ts">
    import CarouselSlideLayout from "@/common-components/Carousel/components/CarouselSlide/CarouselSlideLayout.svelte";

    type CarouselSlideProps = {
        slide: string;
    };

    const {slide}: CarouselSlideProps = $props();
</script>

<CarouselSlideLayout>
    <img
            class="w-full h-auto"
            src={slide}
            alt="Slide"
    />
</CarouselSlideLayout>