export const handleMoreReviewsButtonClick = async (
    setLoading: (loading: boolean) => void,
    showMoreReviewsButtonAction: () => Promise<void>,
) => {
    setLoading(true);
    await showMoreReviewsButtonAction();
    setLoading(false);
}