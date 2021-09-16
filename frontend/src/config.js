
const prod = {
  API_URL: 'https://api-video.convert.p2s.online',
}
 
const dev = {
  API_URL: 'http://localhost:8081',
}

export const config = process.env.NODE_ENV === 'development' ? dev : prod;
