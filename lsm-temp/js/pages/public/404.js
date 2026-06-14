// 404 Error Page - Simple version
import { loadHTMLPage } from "../../helpers/helpers.js";

export default function Error404Page() {
  return {
    async render() {
      return await loadHTMLPage("public", "404");
    },
  };
}