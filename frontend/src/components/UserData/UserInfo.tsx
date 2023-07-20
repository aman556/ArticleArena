import styled from "styled-components";

const UserInfo: React.FC = () => {
  const user_data = {
    name: "Kartikay Sharma",
    email: "kartikaysharma@gmail.com",
    phone: "9876543210",
    Mobile: "9874563210",
    Address: "Saanpo Ka Bill",
  };

  return (
    <Wrapper>
      <table cellSpacing={'10rem'} >
        <tr>
          <th>Full Name</th>
          <td>{user_data.name}</td>
        </tr>
        <hr className="hr"/>
        <tr>
          <th>Email</th>
          <td>{user_data.email}</td>
        </tr>
        <hr className="hr"/>
        <tr>
          <th>Phone</th>
          <td>{user_data.phone}</td>
        </tr>
        <hr className="hr"/>
        <tr>
          <th>Mobile</th>
          <td>{user_data.Mobile}</td>
        </tr>
        <hr className="hr"/>
        <tr>
          <th>Address</th>
          <td>{user_data.Address}</td>
        </tr>
      </table>
    </Wrapper>
  );
};

const Wrapper = styled.section`
  display: flex;
  flex-direction: column;
  gap: 2rem;
  .hr{
     width: 900%;
  },
`;

export default UserInfo;
