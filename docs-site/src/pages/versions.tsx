import type {ReactNode} from 'react';
import Layout from '@theme/Layout';
import Heading from '@theme/Heading';
import Link from '@docusaurus/Link';
import {useEffect, useState} from 'react';
import styles from './versions.module.css';

interface Version {
  version: string;
  date: string;
  current?: boolean;
}

export default function Versions(): ReactNode {
  const [versions, setVersions] = useState<Version[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch('/versions.json')
      .then(res => res.json())
      .then(data => {
        setVersions(data.versions);
        setLoading(false);
      })
      .catch(() => {
        setLoading(false);
      });
  }, []);

  return (
    <Layout
      title="Documentation Versions"
      description="Access different versions of the Bar Out Adapters documentation">
      <main className={styles.container}>
        <div className="container">
          <Heading as="h1">Documentation Versions</Heading>
          
          <div className={styles.intro}>
            <p>
              Select a version below to access the corresponding documentation.
            </p>

            {loading ? (
              <p>Loading versions...</p>
            ) : versions.length > 0 ? (
              <div className={styles.versionsList}>
                {versions.map((v) => (
                  <div
                    key={v.version}
                    className={styles.versionItem}>
                    <div className={styles.versionInfo}>
                      <strong className={styles.versionTitle}>
                        {v.version}
                        {v.current && (
                          <span className={styles.badge}>
                            Current
                          </span>
                        )}
                      </strong>
                      <p className={styles.versionDate}>
                        Released: {v.date}
                      </p>
                    </div>
                    <Link
                      to={v.current ? '/docs/intro' : `/versions/${v.version}/docs/intro`}
                      className="button button--primary">
                      View Docs
                    </Link>
                  </div>
                ))}
              </div>
            ) : (
              <div className={styles.emptyState}>
                <p>
                  No versioned releases yet. View the{' '}
                  <Link to="/docs/intro">current documentation</Link>.
                </p>
              </div>
            )}
          </div>

          <div className={styles.developmentSection}>
            <Heading as="h2" className={styles.sectionTitle}>
              Current Development
            </Heading>
            <p>
              The main documentation is always available at{' '}
              <Link to="/docs/intro">this site</Link> and reflects the latest development
              version.
            </p>
          </div>
        </div>
      </main>
    </Layout>
  );
}
