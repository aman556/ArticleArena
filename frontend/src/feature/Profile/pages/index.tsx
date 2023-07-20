import React from 'react';
import UserCard from '../components/UserCard';
import useStyles from './styles';
import styled from "styled-components";
import Header from '../../Header/Header';
import UserInfo from '../../../components/UserData/UserInfo';

const Profile: React.FC = () => {
  const styles = useStyles();

  return (
    <div>
      <Header />
      <Wrapper>
        <div >
          <div className="grid">
            <div className='container-2'>
              <div className='container-colum-2'>
                <UserCard
                  name={'Kartikay Sharma'}
                  address={'San Francisco, CA'}
                  designation={'Full Stack Developer'}
                />
              </div>
              <div className='container-colum-2'>
                <div>
                  <UserInfo />
                </div>
              </div>

            </div>
            <div className='container-2'>
              <div className='container-colum-2'>
                <div>container 3</div>
              </div>
              <div className='container-colum-2'>
                <div>container 4</div>
              </div>
            </div>
          </div>
        </div>
      </Wrapper>
    </div>
  );
};

const Wrapper = styled.section`
  padding: 0.5rem;

  .grid {
    gap: 4.rem;
  }

  .container-1,
  .container-2,
  .container-3 {
    width: auto;
    height: auto;
    display: flex;
    flex-direction: row;
    background: ${({ theme }) => theme.colors.bg};
    border-radius: 1rem;
    box-shadow: rgba(0, 0, 0, 0.05) 0px 1px 2px 0px;
    padding: 1rem;
    justify-content: center;
    align-content: center;
  }

  .container-2 {
    gap: 2rem;
    background-color: transparent;
    box-shadow: none;

    .container-colum-2 {
      background: ${({ theme }) => theme.colors.bg};
      display: flex;
      flex-direction: row;
      flex: 1;
      border-radius: 0.5rem;
      box-shadow: rgba(0, 0, 0, 0.05) 0px 1px 2px 0px;
      justify-content: auto;
      align-items: auto;

      div {
        display: flex;
        flex-direction: auto;
        gap: 1rem;
        justify-content: center;
        align-items: center;
      }
    }
  }

  h3 {
    margin-top: 1.4rem;
    font-size: 2rem;
  }

  .icon {
    /* font-size: rem; */
    width: 8rem;
    height: 8rem;
    padding: 2rem;
    border-radius: 50%;
    background-color: #fff;
    color: #5138ee;
  }
`;

export default Profile;
