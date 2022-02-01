import http from "../http-common";

class YoutubeDataService {
  getAll() {
    return http.get("/sentiment");
  }

  get(id: any) {
    return http.get(`/sentiment/${id}`);
  }

  create(data: any) {
    return http.post("/sentiment", data);
  }
}

export default new YoutubeDataService();