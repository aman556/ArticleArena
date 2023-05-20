interface ArticleItem {
  ArticleTitle: string;
  ArtilceLink: string;
}

export interface UserArticles {
  ArticleCount: number;
  ArticleData: Array<ArticleItem>;
}

export interface UserHandle {
  Name: string;
  Articles: UserArticles;
}

export interface UserProfileArticleHandles {
  Data: Array<UserHandle>;
}

// [
//   {
//     Name: "gfg",
//     Articles: {
//       ArticleCount: 17,
//       ArticleData: [
//         {
//           ArticleTitle: "Mia Khalifa",
//           ArticleLink: "http/pornhub.com",
//         },
//       ],
//     },
//   },
// ];
