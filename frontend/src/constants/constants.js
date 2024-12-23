export const constants = {
    developmentMode: false,
};

export const enableDevelopmentMode = () => {
    constants.developmentMode = true;
    console.warn('development mode enabled');
}
