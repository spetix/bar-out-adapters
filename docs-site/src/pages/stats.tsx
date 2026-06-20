import type {ReactNode} from 'react';
import Layout from '@theme/Layout';
import Heading from '@theme/Heading';
import {useEffect, useState} from 'react';
import styles from './stats.module.css';

interface PackageCoverage {
  package: string;
  coverage: string;
}

interface ProjectStats {
  packagesTested: number;
  packagesPassed: number;
  packagesFailed: number;
  packagesSkipped: number;
  coverageReportUrl: string;
  rawTestResultsUrl: string;
  generatedAt: string;
  packageCoverages: PackageCoverage[];
}

export default function Stats(): ReactNode {
  const [stats, setStats] = useState<ProjectStats | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch('/project-stats.json')
      .then((res) => {
        if (!res.ok) {
          throw new Error('Stats not available');
        }
        return res.json();
      })
      .then((data) => {
        setStats(data);
      })
      .catch(() => {
        setStats(null);
      })
      .finally(() => {
        setLoading(false);
      });
  }, []);

  return (
    <Layout
      title="Project Stats"
      description="Project test and coverage statistics for Bar Out Adapters">
      <main className={styles.container}>
        <div className="container">
          <Heading as="h1">Project Statistics</Heading>
          <p className={styles.subtitle}>
            Access the latest test and coverage results for this project.
          </p>

          <div className={styles.actions}>
            <a className="button button--secondary" href="/coverage/index.html">
              View Coverage Report
            </a>
            <a className="button button--secondary" href="/test-results.json">
              Download Test Results
            </a>
          </div>

          {loading ? (
            <p className={styles.loading}>Loading stats...</p>
          ) : stats ? (
            <>
              <div className={styles.summaryGrid}>
                <div className={styles.statCard}>
                  <span className={styles.statLabel}>Packages tested</span>
                  <strong>{stats.packagesTested}</strong>
                </div>
                <div className={styles.statCard}>
                  <span className={styles.statLabel}>Passed</span>
                  <strong>{stats.packagesPassed}</strong>
                </div>
                <div className={styles.statCard}>
                  <span className={styles.statLabel}>Failed</span>
                  <strong>{stats.packagesFailed}</strong>
                </div>
                <div className={styles.statCard}>
                  <span className={styles.statLabel}>Skipped</span>
                  <strong>{stats.packagesSkipped}</strong>
                </div>
              </div>

              <div className={styles.tableContainer}>
                <Heading as="h2">Package coverage</Heading>
                {stats.packageCoverages.length > 0 ? (
                  <table className={styles.coverageTable}>
                    <thead>
                      <tr>
                        <th>Package</th>
                        <th>Coverage</th>
                      </tr>
                    </thead>
                    <tbody>
                      {stats.packageCoverages.map((item) => (
                        <tr key={item.package}>
                          <td>{item.package}</td>
                          <td>{item.coverage}</td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
                ) : (
                  <p>No coverage data available.</p>
                )}
              </div>

              <div className={styles.generatedAt}>
                Generated: {stats.generatedAt}
              </div>
            </>
          ) : (
            <div className={styles.fallback}>
              <p>Project stats are not yet available. Check back after the next docs build.</p>
            </div>
          )}
        </div>
      </main>
    </Layout>
  );
}
