<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import CarouselLayout from "@/common-components/Carousel/CarouselLayout.svelte";
    import CarouselSlideList from "@/common-components/Carousel/components/CarouselSlideList/CarouselSlideList.svelte";
    import CarouselControl from "@/common-components/Carousel/components/CarouselControl/CarouselControl.svelte";

    type CarouselProps = LayoutProps & {
        slideNumber: number,
    };

    const {children, slideNumber}: CarouselProps = $props();

    let currentIndex = $state(0);

    const nextSlide = () => {
        currentIndex = (currentIndex + 1) % slideNumber;
    };

    const prevSlide = () => {
        currentIndex = (currentIndex - 1 + slideNumber) % slideNumber;
    };
</script>

<CarouselLayout>
    <CarouselSlideList currentIndex={currentIndex}>
        {@render children()}
    </CarouselSlideList>
    <CarouselControl nextSlide={nextSlide} prevSlide={prevSlide}/>
</CarouselLayout>
