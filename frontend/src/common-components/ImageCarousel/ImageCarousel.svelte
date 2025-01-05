<script lang="ts">
    import CarouselLayout from "@/common-components/ImageCarousel/ImageCarouselLayout.svelte";
    import CarouselControl from "@/common-components/ImageCarousel/components/ImageCarouselControl/ImageCarouselControl.svelte";
    import CarouselSlideList from "@/common-components/ImageCarousel/components/ImageCarouselSlideList/ImageCarouselSlideList.svelte";

    type CarouselProps = {
        slides: string[];
    };

    const {slides}: CarouselProps = $props();

    let currentIndex = $state(0);

    const nextSlide = () => {
        currentIndex = (currentIndex + 1) % slides.length;
    };

    const prevSlide = () => {
        currentIndex = (currentIndex - 1 + slides.length) % slides.length;
    };
</script>

<CarouselLayout>
    <CarouselSlideList currentIndex={currentIndex} slides={slides}/>
    <CarouselControl nextSlide={nextSlide} prevSlide={prevSlide}/>
</CarouselLayout>
