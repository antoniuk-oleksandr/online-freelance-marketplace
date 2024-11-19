import {Carousel} from "@mantine/carousel";
import '@mantine/carousel/styles.css';
import {Service} from "@/types/Service";
import ServiceImageCarouselSlide from "../ServiceImageCarouselSlide/ServiceImageCarouselSlide";
import {useTheme} from "next-themes";

type ServiceImageCarouselProps = Service;

const ServiceImageCarousel = (props: ServiceImageCarouselProps) => {
    const {images} = props;
    const {resolvedTheme} = useTheme();
    const isDark = resolvedTheme === "dark";

    return (
        <Carousel styles={{
            control: {
                borderColor: isDark ? 'rgba(255, 255, 255, 0.12)' : 'rgba(0, 0, 0, 0.12)',
                color: isDark ? '#fff' : '#000',
                background: isDark ? '#1a1a1a' : '#f9f9f9',
            },
        }} loop withIndicators>
            {images.map((item, index) => (
                <ServiceImageCarouselSlide key={index} src={item}/>
            ))}
        </Carousel>
    )
}

export default ServiceImageCarousel;