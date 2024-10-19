export default {
  "/index.html": function (isBuild) {
    if (isBuild) {
      return {
        title: "{{.Title}}",
        sections: ["{{.SectionTitle}}"]
      };
    }
    return {
      title: "example title",
      sections: ["Section A", "Section B", "Section C"],
    };
  },
};
