<script lang="ts">
    import CarouselLayout from "@/common-components/Carousel/CarouselLayout.svelte";
    import CarouselControl from "@/common-components/Carousel/components/CarouselControl/CarouselControl.svelte";
    import CarouselSlideList from "@/common-components/Carousel/components/CarouselSlideList/CarouselSlideList.svelte";

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
