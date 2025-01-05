<script lang="ts">
    import CarouselSlideLayout from "@/common-components/ImageCarousel/components/ImageCarouselSlide/ImageCarouselSlideLayout.svelte";

    type CarouselSlideProps = {
        slide: string;
    };

    const {slide}: CarouselSlideProps = $props();
</script>

<CarouselSlideLayout>
    <img
            class="aspect-video size-full object-cover object-center"
            src={slide}
            alt="Slide"
    />
</CarouselSlideLayout>